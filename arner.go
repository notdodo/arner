package arner

import (
	"encoding/json"
	"errors"
	"strings"
)

const (
	arnDelimiter = ":"
	arnSections  = 6
	arnPrefix    = "arn:"

	// zero-indexed
	sectionPartition = 1
	sectionService   = 2
	sectionRegion    = 3
	sectionAccountID = 4
	sectionResource  = 5

	// errors
	invalidPrefix = "arn: invalid prefix"
)

// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/aws/arn
// https://github.com/aws/aws-sdk-go-v2/blob/v1.16.11/aws/arn/arn.go#L2
type BetterARN struct {
	// The partition that the resource is in. For standard AWS regions, the partition is "aws". If you have resources in
	// other partitions, the partition is "aws-partitionname". For example, the partition for resources in the China
	// (Beijing) region is "aws-cn".
	Partition string

	// The service namespace that identifies the AWS product (for example, Amazon S3, IAM, or Amazon RDS). For a list of
	// namespaces, see
	// http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces.
	Service string

	// The region the resource resides in. Note that the ARNs for some resources do not require a region, so this
	// component might be omitted.
	Region string

	// The ID of the AWS account that owns the resource, without the hyphens. For example, 123456789012. Note that the
	// ARNs for some resources don't require an account number, so this component might be omitted.
	AccountID string

	// The content of this part of the ARN varies by service. It often includes an indicator of the type of resource —
	// for example, an IAM user or Amazon RDS database - followed by a slash (/) or a colon (:), followed by the
	// resource name itself. Some services allows paths for resource names, as described in
	// http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arns-paths.
	ResourceType string

	// The actual resource name that the ARN is pointing to;
	// for example, a DynamoDB table arn:aws:dynamodb:::table/thistableprod real name is thistableprod.
	Resource string

	// Paths can be use to logically separate resources like Organization Units for Users/Roles or specify subdirectory in S3 buckets.
	Path string
}

func ParseARN(arn string) (BetterARN, error) {
	if !IsARN(arn) {
		return BetterARN{}, errors.New(invalidPrefix)
	}

	var (
		resourceType string
		resourceName string
		path         string
	)
	sections := strings.SplitN(arn, arnDelimiter, arnSections)

	v := sections[sectionService]
	switch {
	case v == "iam":
		resourceType, resourceName, path = parseSlash(sections[sectionResource])
	case v == "s3":
		// S3 ARNs are "special": https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-arn-format.html
		resourceName, path = parseS3(sections[sectionResource])
	case v == "dynamodb":
		resourceType, resourceName, path = parseSlash(sections[sectionResource])
	case v == "ec2":
		resourceType, resourceName, path = parseSlash(sections[sectionResource])
	case v == "lambda":
		resourceType, resourceName = parseColon(sections[sectionResource])
	case v == "rds":
		resourceType, resourceName = parseColon(sections[sectionResource])
	case v == "redshift":
		resourceType, resourceName = parseColon(sections[sectionResource])
	default:
		resourceName = sections[sectionResource]
	}

	return BetterARN{
		Partition:    sections[sectionPartition],
		Service:      sections[sectionService],
		Region:       sections[sectionRegion],
		AccountID:    sections[sectionAccountID],
		Resource:     resourceName,
		ResourceType: resourceType,
		Path:         path,
	}, nil
}

func parseSlash(resource string) (resType, resName, resPath string) {
	pathSplits := strings.Split(resource, "/")
	if len(pathSplits) > 2 {
		resPath = "/" + strings.Join(pathSplits[1:len(pathSplits)-1], "/") + "/"
	}
	return pathSplits[0], pathSplits[len(pathSplits)-1], resPath
}

func parseS3(resource string) (bucketName, resPath string) {
	pathSplits := strings.Split(resource, "/")
	if len(pathSplits) > 1 {
		if pathSplits[len(pathSplits)-1] == "*" {
			resPath = strings.Join(pathSplits[1:len(pathSplits)-1], "/")
		} else {
			resPath = strings.Join(pathSplits[1:], "/")
		}
	}
	return pathSplits[0], resPath
}

func parseColon(resource string) (resourceType, funName string) {
	pathSplits := strings.Split(resource, ":")
	return pathSplits[0], pathSplits[len(pathSplits)-1]
}

// IsARN returns whether the given string is an arn
// by looking for whether the string starts with arn:
func IsARN(arn string) bool {
	return strings.HasPrefix(arn, arnPrefix) && strings.Count(arn, ":") >= arnSections-1
}

// String returns the canonical representation of the ARN
func (arn BetterARN) String() string {
	arnString, _ := json.Marshal(arn)
	return string(arnString)
}
