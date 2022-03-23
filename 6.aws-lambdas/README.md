# Lambdas in AWS

## create a S3 bucket
- [doc](https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html) The easiest way to create a S3 bucket is to use the UI.

- put a text and a photo into the S3 bucket

Note: Try not to dive too deep into AWS, as it quickly drifts you from the focus of this exercise which focuses more on the golang backend.

## Download and install AWS sdk go
- [doc]()

## Learn to use AWS serverless: Lambda
- [doc]()

---
Good for future learning

## AWS S3 Bucket Service:
understand the key in S3 bucket
### 1. Versioning
- how to turn on versioning for a bucket
- how to restore a deleted file
- how to permanently delete a file
### 2. Encryption
4 methods to encrypt objects in S3, SSE means server side encryption
- SSE-S3: encrypts S3 objects using keys. Keys are handled and managed by AWS
- SSE-KMS: encryption using keys handled and managed by KMS (user control + audit trail)
- SSE-C: encryption using keys outside of AWS and are fully managed by customers, must use HTTPS
- client side encryption: client will perform the encryption on the client side and upload the encrypted object to S3 bucket

Finally Encryption in transit(SSL/TLS):
Amazon S3 exposes:
- HTTP endpoint
- HTTPS endpoint: encryption in flight
### 3. Bucket Policies and Security
- User based: IAM Policies
- Resource based:
  1. Bucket Policies
  2. Object ACL - finer grain
  3. Bucket ACL - less common
### 4. S3 Websites
- S3 can host static website 
- the website url will be:
<bucket-name>.s3-website<AWS-region>.amazonaws.com
- make sure the bucket policy allows the public reads, otherwise you will get the 403 forbidden error

### 5. S3 CORS: cross-origin request 

## Lambda
1. Virtual functions
2. Limited by time - short executions
3. Run on-demand
4. scaling is automated
