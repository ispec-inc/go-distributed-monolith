provider "aws" {
  region  = "us-east-1"
  version = "3.50.0"
  alias   = "virginia"
}

data "archive_file" "file" {
  type        = "zip"
  source_dir  = "${var.src_path}/src"
  output_path = "${var.src_path}/dist/package.zip"
}

data "aws_partition" "current" {}

data "aws_iam_policy_document" "execution_role" {
  statement {
    sid = "AllowCloudWatchLogs"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    effect = "Allow"
    resources = [
      format(
        "arn:%s:logs:*::log-group:/aws/lambda/*:*:*",
        data.aws_partition.current.partition
      )
    ]
  }
}

resource "aws_iam_role" "execution_role" {
  name               = "${var.env}-${var.role}"
  assume_role_policy = <<-EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Principal": {
          "Service": [
            "lambda.amazonaws.com",
            "edgelambda.amazonaws.com"
          ]
        },
        "Effect": "Allow",
        "Sid": ""
      }
    ]
  }
  EOF
}

resource "aws_iam_policy" "execution_role" {
  name   = "${var.env}-${var.role}"
  path   = "/"
  policy = data.aws_iam_policy_document.execution_role.json
}

resource "aws_iam_role_policy_attachment" "execution_role" {
  role       = aws_iam_role.execution_role.name
  policy_arn = aws_iam_policy.execution_role.arn
}

resource "aws_lambda_function" "lambda" {
  provider         = aws.virginia
  filename         = data.archive_file.file.output_path
  function_name    = "${var.env}-${var.role}"
  role             = aws_iam_role.execution_role.arn
  handler          = var.handler
  source_code_hash = data.archive_file.file.output_base64sha256
  runtime          = var.runtime
  publish          = var.publish
  memory_size      = var.memory_size
  timeout          = var.timeout
}
