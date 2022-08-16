# ARNer

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-93%25-brightgreen.svg?longCache=true&style=flat)</a>

> Amazon Resource Names (ARNs) uniquely identify AWS resources. We require an ARN when you need to specify a resource unambiguously across all of AWS, such as in IAM policies, Amazon Relational Database Service (Amazon RDS) tags, and API calls.

[Amazon Resource Names Documentation](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)

The official AWS Golang SDK library has a simple ARN parsing function ([Parse](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2@v1.16.11/aws/arn#Parse)) which is not actually accurate if you need to know the exact resource's name and/or path and/or resource type (i.e. role, group, table, cluster, etc.)

This library correctly parse AWS ARN strings and outputs a BetterARN struct.

```golang
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
```

## Example

```golang
package main

import (
	"fmt"
	"log"

	"github.com/notdodo/arner"
)

func main() {
	arn := "arn:aws:iam::111111111111:user/division_abc/subdivision_xyz/JaneDoe"

	betterARN, err := arner.ParseARN(arn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(betterARN)
}
```

Output will be:

```json
{
  "Partition": "aws",
  "Service": "iam",
  "Region": "",
  "AccountID": "111111111111",
  "ResourceType": "user",
  "Resource": "JaneDoe",
  "Path": "division_abc/subdivision_xyz"
}
```

## Why?

Because the default AWS SDK Golang library V2, for the previous example returns:

```json
{
  "Partition": "aws",
  "Service": "iam",
  "Region": "",
  "AccountID": "111111111111",
  "Resource": "user/division_abc/subdivision_xyz/JaneDoe"
}
```

In my case I need the exact name of the resource, not the complete path plus the resource type plus other garbage.
