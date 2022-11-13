resource "aws_dynamodb_table" "capture" {
  name         = "inc${terraform.workspace}capture"
  hash_key     = "network"
  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "network"
    type = "S"
  }
}

resource "aws_dynamodb_table" "safe" {
  name         = "inc${terraform.workspace}safe"
  hash_key     = "address"
  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "address"
    type = "S"
  }
}

resource "aws_dynamodb_table" "signer" {
  name         = "inc${terraform.workspace}signer"
  hash_key     = "address"
  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "address"
    type = "S"
  }

}

data "aws_iam_policy_document" "authorization" {
  policy_id = "inc${terraform.workspace}dynamo"
  statement {
    sid    = "incdynamo"
    effect = "Allow"
    actions = [
      "dynamodb:PutItem",
      "dynamodb:DeleteItem",
      "dynamodb:GetItem",
    ]
    resources = [
      aws_dynamodb_table.capture.arn,
      aws_dynamodb_table.safe.arn,
      aws_dynamodb_table.signer.arn,
    ]
  }
}

resource "aws_iam_policy" "dynamodb_policy" {
  name        = "inc${terraform.workspace}dynamopolicy"
  description = "The policy to access the INC tables"
  policy      = data.aws_iam_policy_document.authorization.json
}

resource "aws_iam_role_policy_attachment" "dynamodb_for_lambda" {
  policy_arn = aws_iam_policy.dynamodb_policy.arn
  role       = aws_iam_role.iam_for_lambda.name
}
