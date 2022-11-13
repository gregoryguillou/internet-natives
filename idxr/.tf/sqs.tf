resource "aws_lambda_event_source_mapping" "sqstrigger" {
  event_source_arn = aws_sqs_queue.sqs_queue_trigger.arn
  enabled          = true
  function_name    = aws_lambda_function.lambda.arn
}

resource "aws_sqs_queue" "sqs_queue_trigger" {
  name                      = "inc${terraform.workspace}trigger"
  max_message_size          = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 3
  redrive_policy            = "{\"deadLetterTargetArn\":\"${aws_sqs_queue.sqs_queue_deadletter.arn}\",\"maxReceiveCount\":4}"

  tags {
    Environment = "production"
  }
}

resource "aws_sqs_queue" "sqs_queue_deadletter" {
  name                      = "inc${terraform.workspace}deadletter"
  max_message_size          = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10

  tags {
    Environment = "production"
  }
}


data "aws_iam_policy_document" "authorization" {
  policy_id = "inc${terraform.workspace}sqs"
  statement {
    sid    = "incsqs"
    effect = "Allow"
    actions = [
      "sqs:ReceiveMessage",
      "sqs:DeleteMessage",
      "sqs:GetQueueAttributes",
    ]
    resources = [
      aws_sqs_queue.sqs_queue_trigger.arn,
    ]
  }
}

resource "aws_iam_policy" "sqs_policy" {
  name        = "inc${terraform.workspace}sqspolicy"
  description = "The policy to access the INC sqs"
  policy      = data.aws_iam_policy_document.authorization.json
}

resource "aws_iam_role_policy_attachment" "dynamodb_for_lambda" {
  policy_arn = aws_iam_policy.dynamodb_policy.arn
  role       = aws_iam_role.iam_for_lambda.name
}
