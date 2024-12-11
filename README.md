# Overview
This project can be seen as a proof of concept, trying to show the functionality of SSI in combination with everyday email communication. It contains two programs:
- Janus: Janus is a mail agent written in C#. Capable of communicating via IMAP/POP with your email provider. With it, you can send and read emails using SSI security.

- Monsi: Monsi is a SSI wallet implementation written in Go using [Wails](https://wails.io/). It includes endpoints for the communication with Janus and manages encryption and decryption of mail contents.

# Setup
There are binaries for immediate testing for Windows systems here:
- .\monsi\build\bin

Linux and Mac will have to build there own solutions.

To run **Janus**, you will need:  
- Visual Studio  
- .NET 6  

### Configure User Secrets  

Before running the project, edit the user secrets and insert the following configuration:  


```json
{
  "EmailSettings": {
    "Email": "myMail",
    "Password": "myPassword"
  },
  "ServerSettings": {
    "SMTPServer": "MyMailProviderSMTPServer",
    "SMTPPort": MyMailProviderSMTPPort,
    "IMAPSever": "MyMailproviderIMAPServer",
    "IMAPPort": MyMailProviderIMAPPort
  }
}
```

Replace the placeholder values with your actual email address, password, SMTP server name, and port provided by your email provider.

For example, if you are using **web.de**, the configuration would look something like this:
```json
{
  "EmailSettings": {
    "Email": "qwertz0014@web.de",
    "Password": "SSITe5tM@ilSaf3"
  },
  "ServerSettings": {
    "SMTPServer": "smtp.web.de",
    "SMTPPort": 587,
    "IMAPSever": "imap.web.de",
    "IMAPPort": 993
  }
}
```

### Run the Project

Once the user secrets are configured, you can:

- Run the project directly in Visual Studio.
- Publish it and host it yourself, for example, on IIS.
  
# Future work
- Encryption via password of DIDs (and especially its passwords) on disk.

# Remarks 
- The logos for Janus and Monsi are AI generated using DALL-E