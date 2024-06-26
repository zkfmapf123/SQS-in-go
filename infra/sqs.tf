################################### SQS Queue (DeadLetter Queue) ############################
resource "aws_sqs_queue" "dead_letter_queue" {
  name = "ex-dead-letter-queue"
}


################################### SQS Queue ############################
resource "aws_sqs_queue" "queue" {
  name = "ex-queue"

  delay_seconds             = 90    ## 메시지 전송 후 대기 시간
  max_message_size          = 2048  ## 최대 메시지 크기
  message_retention_seconds = 86400 ## 메시지 보존 기간 1일 (60 * 60 * 24)
  receive_wait_time_seconds = 10    ## 메시지 수신 대기 시간

  ## 메시지 재전송 정책
  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.dead_letter_queue.arn
    maxReceiveCount     = 5 ## 총 4번 시도 한 후 Dead Letter Queue에 쌓임
  })

  depends_on = [aws_sqs_queue.dead_letter_queue]

  tags = {
    Environment = "test"
  }
}

################################### SQS Queue FIFO ############################
## FIFO Queue는 DeadLetter 큐를 사용할 수 없음
resource "aws_sqs_queue" "fifo-queue" {
  name                        = "ex-queue.fifo"
  fifo_queue                  = true
  content_based_deduplication = true ## 메시지 중복 방지 (여러번 쌓여도 중복 제거됨)

  depends_on = [aws_sqs_queue.dead_letter_queue]

  tags = {
    Environment = "test"
  }
}


output "queue_url" {
  value = {
    dead_queue = aws_sqs_queue.dead_letter_queue.url
    queue      = aws_sqs_queue.queue.url
    fifo_queue = aws_sqs_queue.fifo-queue.url
  }
}
