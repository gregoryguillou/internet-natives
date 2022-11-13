data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "iam_for_lambda" {
  name               = "inc${terraform.workspace}api"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
}

resource "aws_iam_role_policy_attachment" "iam_for_lambda" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.iam_for_lambda.name
}

resource "aws_lambda_function" "lambda" {
  s3_bucket     = var.bucket
  s3_key        = "lambda/api/${data.external.version.result["version"]}.zip"
  function_name = "inc${terraform.workspace}api"
  role          = aws_iam_role.iam_for_lambda.arn
  runtime       = "go1.x"
  memory_size   = 512
  timeout       = 30
  handler       = "api-linux-amd64"

  environment {
    variables = {
      route_prefix = "/production"
      version      = data.external.version.result["version"]
    }
  }
}