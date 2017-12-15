package service

// imports have been intentionally removed

type S3Service struct {
	
}

func (s3s *S3Service) UploadFileToS3(destinationPath string, sourcePath string) (error, string) {

	// The session the S3 Uploader will use
	session := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(session)

	f, err := os.Open(sourcePath)
	if err != nil {
		fmt.Errorf("failed to open file %q, %v", sourcePath, err)
		return err, ""
	}

	defer f.Close()

	fmt.Print("s3 bucket name : " + cmlconstants.S3BucketName + destinationPath + "\n")
	fmt.Print("s3 key : " + filepath.Base(sourcePath) + "\n")

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cmlconstants.S3BucketName + destinationPath),
		Key:    aws.String(filepath.Base(sourcePath)),
		Body:   f,
	})

	if err != nil {
		fmt.Print(err.Error())
		fmt.Errorf("failed to upload file %q, %v", filepath.Base(sourcePath), err)
		return err, ""
	}

	fmt.Print("upload location ", result.Location)
	return nil, destinationPath + filepath.Base(sourcePath)
}

func (s3s *S3Service) DeleteS3Object(bucket string, key string) error {
	// Initialize a session that the SDK uses to load configuration,
	// credentials, and region from the shared config file. (~/.aws/config).
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create S3 service client
	svc := s3.New(session)

	// Delete the item
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)})

	if err != nil {
		fmt.Print("s3 object delete error : " + err.Error())
		return err
	}

	return err
}

func (s3s *S3Service) CheckIfObjectExists(bucket string, key string) error {
	// Initialize a session that the SDK uses to load configuration,
	// credentials, and region from the shared config file. (~/.aws/config).
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create S3 service client
	svc := s3.New(session)

	headInput := s3.HeadObjectInput{Bucket: &bucket, Key: &key}

	output, err := svc.HeadObject(&headInput)
	fmt.Print(output)
	if err != nil {
		return err
	}
	return nil
}
