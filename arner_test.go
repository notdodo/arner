package arner

import (
	"reflect"
	"testing"
)

var ARNs = []struct {
	test     string
	expected BetterARN
}{
	{
		test: "arn:aws:iam::123456789012:user/JohnDoe",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "user",
			Resource:     "JohnDoe",
			Path:         "",
		},
	},
	{
		test: "arn:aws:iam::123456789012:user/division_abc/subdivision_xyz/JaneDoe",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "user",
			Resource:     "JaneDoe",
			Path:         "division_abc/subdivision_xyz",
		},
	},
	{
		test: "arn:aws:iam::123456789012:user/division_abc/JaneDoe",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "user",
			Resource:     "JaneDoe",
			Path:         "division_abc",
		},
	},
	{
		test: "arn:aws:iam::123456789012:group/Developers",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "group",
			Resource:     "Developers",
			Path:         "",
		},
	},
	{
		test: "arn:aws:iam::123456789012:group/division_abc/subdivision_xyz/product_A/Developers",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "group",
			Resource:     "Developers",
			Path:         "division_abc/subdivision_xyz/product_A",
		},
	},
	{
		test: "arn:aws:iam::123456789012:role/S3Access",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "role",
			Resource:     "S3Access",
			Path:         "",
		},
	},
	{
		test: "arn:aws:iam::123456789012:role/application_abc/component_xyz/RDSAccess",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "role",
			Resource:     "RDSAccess",
			Path:         "application_abc/component_xyz",
		},
	},
	{
		test: "arn:aws:iam::123456789012:role/aws-service-role/access-analyzer.amazonaws.com/AWSServiceRoleForAccessAnalyzer",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "role",
			Resource:     "AWSServiceRoleForAccessAnalyzer",
			Path:         "aws-service-role/access-analyzer.amazonaws.com",
		},
	},
	{
		test: "arn:aws:iam::123456789012:role/service-role/QuickSightAction",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "role",
			Resource:     "QuickSightAction",
			Path:         "service-role",
		},
	},
	{
		test: "arn:aws:iam::123456789012:policy/UsersManageOwnCredentials",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "policy",
			Resource:     "UsersManageOwnCredentials",
			Path:         "",
		},
	},
	{
		test: "arn:aws:iam::123456789012:policy/division_abc/subdivision_xyz/UsersManageOwnCredentials",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "policy",
			Resource:     "UsersManageOwnCredentials",
			Path:         "division_abc/subdivision_xyz",
		},
	},
	{
		test: "arn:aws:iam::123456789012:instance-profile/Webserver",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "instance-profile",
			Resource:     "Webserver",
			Path:         "",
		},
	},
	{
		test: "arn:aws:sts::123456789012:federated-user/JohnDoe",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "sts",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "",
			Resource:     "federated-user/JohnDoe",
			Path:         "",
		},
	},
	{
		test: "arn:aws:sts::123456789012:assumed-role/Accounting-Role/JaneDoe",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "sts",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "",
			Resource:     "assumed-role/Accounting-Role/JaneDoe",
			Path:         "",
		},
	},
	{
		test: "arn:aws:iam::123456789012:mfa/JaneDoeMFA",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "mfa",
			Resource:     "JaneDoeMFA",
			Path:         "",
		},
	},
	{
		test: "arn:aws:iam::123456789012:u2f/user/JohnDoe/default (U2F security key)",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "u2f",
			Resource:     "default (U2F security key)",
			Path:         "user/JohnDoe",
		},
	},
	{
		test: "arn:aws:iam::123456789012:server-certificate/ProdServerCert",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "server-certificate",
			Resource:     "ProdServerCert",
			Path:         "",
		},
	},
	{
		test: "arn:aws:iam::123456789012:server-certificate/division_abc/subdivision_xyz/ProdServerCert",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "server-certificate",
			Resource:     "ProdServerCert",
			Path:         "division_abc/subdivision_xyz",
		},
	},
	{
		test: "arn:aws:iam::123456789012:saml-provider/ADFSProvider",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "saml-provider",
			Resource:     "ADFSProvider",
			Path:         "",
		},
	},
	{
		test: "arn:aws:iam::123456789012:oidc-provider/GoogleProvider",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "oidc-provider",
			Resource:     "GoogleProvider",
			Path:         "",
		},
	},
	{
		test:     "arn:garbage",
		expected: BetterARN{},
	},
	{
		test:     "more:Garbage::::junk",
		expected: BetterARN{}},
	{
		test: "arn:aws:iam::123456789012:instance-profile/actual-instance-profile-role-name",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "instance-profile",
			Resource:     "actual-instance-profile-role-name",
			Path:         "",
		},
	},
	{
		test: "arn:aws:iam::123456789012:instance-profile/division_abc/actual-instance-profile-role-name",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "instance-profile",
			Resource:     "actual-instance-profile-role-name",
			Path:         "division_abc",
		},
	},
	{
		test: "arn:aws:iam::123456789012:user/actual-user-name",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "user",
			Resource:     "actual-user-name",
			Path:         "",
		},
	},
	{
		test: "arn:aws:iam::123456789012:group/actual-group-name",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "group",
			Resource:     "actual-group-name",
			Path:         "",
		},
	},
	{
		test: "arn:aws:iam::123456789012:role/aws-service-role/ops.apigateway.amazonaws.com/AWSServiceRoleForAPIGateway",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "iam",
			Region:       "",
			AccountID:    "123456789012",
			ResourceType: "role",
			Resource:     "AWSServiceRoleForAPIGateway",
			Path:         "aws-service-role/ops.apigateway.amazonaws.com",
		},
	},
	{
		test: "arn:aws:s3:::bucket-name/*",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "s3",
			Region:       "",
			AccountID:    "",
			ResourceType: "",
			Resource:     "bucket-name",
			Path:         "",
		},
	},
	{
		test: "arn:aws:s3:::bucket-name",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "s3",
			Region:       "",
			AccountID:    "",
			ResourceType: "",
			Resource:     "bucket-name",
			Path:         "",
		},
	},
	{
		test: "arn:aws:s3:::bucket-name/bucket-folder",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "s3",
			Region:       "",
			AccountID:    "",
			ResourceType: "",
			Resource:     "bucket-name",
			Path:         "",
		},
	},
	{
		test: "arn:aws:s3:::bucket-name/bucket-folder/*",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "s3",
			Region:       "",
			AccountID:    "",
			ResourceType: "",
			Resource:     "bucket-name",
			Path:         "bucket-folder",
		},
	},
	{
		test: "arn:aws:dynamodb:eu-west-1:123456789012:table/actual-table-name",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "dynamodb",
			Region:       "eu-west-1",
			AccountID:    "123456789012",
			ResourceType: "table",
			Resource:     "actual-table-name",
			Path:         "",
		},
	},
	{
		test: "arn:aws:ec2:ap-south-1:123456789012:instance/i-11b11b11a1a111111",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "ec2",
			Region:       "ap-south-1",
			AccountID:    "123456789012",
			ResourceType: "instance",
			Resource:     "i-11b11b11a1a111111",
			Path:         "",
		},
	},
	{
		test: "arn:aws:lambda:eu-west-1:123456789012:function:lambda-name",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "lambda",
			Region:       "eu-west-1",
			AccountID:    "123456789012",
			ResourceType: "function",
			Resource:     "lambda-name",
			Path:         "",
		},
	},
	{
		test: "arn:aws:rds:eu-west-1:123456789012:cluster:rds-cluster-name",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "rds",
			Region:       "eu-west-1",
			AccountID:    "123456789012",
			ResourceType: "cluster",
			Resource:     "rds-cluster-name",
			Path:         "",
		},
	},
	{
		test: "arn:aws:rds:eu-west-1:123456789012:db:rds-instance-name",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "rds",
			Region:       "eu-west-1",
			AccountID:    "123456789012",
			ResourceType: "db",
			Resource:     "rds-instance-name",
			Path:         "",
		},
	},
	{
		test: "arn:aws:redshift:eu-west-1:123456789012:namespace:1c11111c-1fc1-1ee1-11bc-111111111a11",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "redshift",
			Region:       "eu-west-1",
			AccountID:    "123456789012",
			ResourceType: "namespace",
			Resource:     "1c11111c-1fc1-1ee1-11bc-111111111a11",
			Path:         "",
		},
	},
	{
		test: "arn:aws:redshift:us-west-2:123456789012:cluster:cluster-name",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "redshift",
			Region:       "us-west-2",
			AccountID:    "123456789012",
			ResourceType: "cluster",
			Resource:     "cluster-name",
			Path:         "",
		},
	},
	{
		test: "arn:aws:ec2:us-east-1:123456789012:vpc/vpc-1234567890abcdef0",
		expected: BetterARN{
			Partition:    "aws",
			Service:      "ec2",
			Region:       "us-east-1",
			AccountID:    "123456789012",
			ResourceType: "vpc",
			Resource:     "vpc-1234567890abcdef0",
			Path:         "",
		},
	},
	{
		test: "arn:*:iam::*:role/AWS-QuickSetup-StackSet-Local-ExecutionRole", // yes, this is real
		expected: BetterARN{
			Partition:    "*",
			Service:      "iam",
			Region:       "",
			AccountID:    "*",
			ResourceType: "role",
			Resource:     "AWS-QuickSetup-StackSet-Local-ExecutionRole",
			Path:         "",
		},
	},
}

func TestSupportedServices(t *testing.T) {
	for _, test := range ARNs {
		got, err := ParseARN(test.test)

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("mismatch!\n\tgot: %v\n\twanted: %v\n\toriginal: %s", got, test.expected, test.test)
			if err != nil {
				t.Errorf("[X] Test failed with error: %v", err)
				continue
			}
		}

	}
}
