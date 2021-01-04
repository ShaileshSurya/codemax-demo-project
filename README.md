# Mail Service Rest API. 

Send an Email by rest API where to, BCC and CC will be provided via api-payload. 

1. Use design pattern
    > Used design pattern factory and strategy. Made sure to code to interface and not the concreate classes. 

2. Use middleware to store the logging
    > Done, can check the loggermiddleware.go file. 

3. You can use any framework or api provider in golang.
    > Used negroni and servemux. No framework as such. 

4. In future we will use Amazon SES and sendgrid, take this part into consideration. Don’t
implement the coding for this but make some support coding so it could be easily
implement sendgrid or Amazon ses or any other providers.
    > There is a factory method that fetches the mailer. The handlers or user have no knowledge of the underlying provider. 

5. Add normal API token authentication static (x-api)
    > Done. Created a middleware check xapiauthenticationmiddleware. 

6. Create git repo and please add the documentation and instruction to deploy the demo in
any local machine.
    > Done. 

    How to deploy this app. Use the following commands. 

    1. git clone https://github.com/ShaileshSurya/mail-service-rest-api-project.git
    2. cd mail-service-rest-api-project
    3. go build
    4. ./mail-service-rest-api-project
    5. server is hardcoded to run on `port 8082`. 
    6. use following curl command to send email.

    ```bash
    
    curl --location --request POST 'http://localhost:8082/mail' \
            --header 'x-api-key: 123456789' \
            --header 'Content-Type: text/plain' \
            --data-raw '{
            "to": [
                {
                "name": "shailesh",
                "email": "suryas3.social@gmail.com"
                }
            ],
            "cc": [
                {
                "name": "shailesh",
                "email": "suryas3.social@gmail.com"
                },
                {
                "name": "shailesh",
                "email": "suryas3.social2@gmail.com"
                }
            ],
            "bcc": [
                {
                "name": "shailesh",
                "email": "suryas3.social@gmail.com"
                },
                {
                "name": "shailesh",
                "email": "suryas3.social4@gmail.com"
                }
            ],
            "subject": "test_mail_sending",
            "body": "sample_body_test_message"
            }'
    
    ```

    Future scope. 
    1. Write UT's. 
    2. Write E2E. 


Some sample log statements. 

```bash
root@ubuntu1604:~/go_project/src/github.com/ShaileshSurya/mail-service-rest-api-project# ./mail-service-rest-api-project
INFO[0000] /mail
INFO[0000] Email sending service RESTfully... on port : 8082
ERRO[0012] Error while unmarshalling data.invalid character '\r' in string literal
INFO[0012] Request Log dump                              duration="276.529µs" method=POST path=/mail status=400
[negroni] 2020-12-19T05:56:29-05:00 | 400 |      460.838µs | localhost:8082 | POST /mail
INFO[0033] Data: &{ResultsV31:[{Status:success CustomID: To:[{Email:suryas3.social@gmail.com MessageUUID:c3708070-5dd1-4ad9-b1f1-96008fac0083 MessageID:1152921510239503004 MessageHref:https://api.mailjet.com/v3/REST/message/1152921510239503004}] Cc:[{Email:suryas3.social@gmail.com MessageUUID:92de88fb-6aa6-40ab-80bb-2311d0dc036c MessageID:1152921510239503005 MessageHref:https://api.mailjet.com/v3/REST/message/1152921510239503005} {Email:suryas3.social2@gmail.com MessageUUID:e56c4f90-8b65-4a45-9fee-e18b146c4c9d MessageID:1152921510239503006 MessageHref:https://api.mailjet.com/v3/REST/message/1152921510239503006}] Bcc:[{Email:suryas3.social@gmail.com MessageUUID:51ee1e05-c06c-4585-bc1e-9dbc3c070cbe MessageID:1152921510239503007 MessageHref:https://api.mailjet.com/v3/REST/message/1152921510239503007} {Email:suryas3.social4@gmail.com MessageUUID:8497b6e1-979a-4c48-9baf-7f56b7eda8c6 MessageID:1152921510239503008 MessageHref:https://api.mailjet.com/v3/REST/message/1152921510239503008}]}]}
INFO[0033] Request Log dump                              duration=666.77805ms method=POST path=/mail status=201
[negroni] 2020-12-19T05:56:49-05:00 | 201 |      666.82343ms | localhost:8082 | POST /mail

```
