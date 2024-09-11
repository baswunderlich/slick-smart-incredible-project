using MailKit.Net.Smtp;
using MimeKit;
using Microsoft.AspNetCore.Mvc;
using JanusWeb.Models;
using System.Text.Json;
using System.Text;
using JanusWeb.Business;

namespace JanusWeb.Controllers
{
    public class MailCreatorController : Controller
    {
        [HttpGet]
        public IActionResult Index()
        {
            return View();  // Return the page to create and send an email
        }

        [HttpPost]
        public async Task<IActionResult> SendEmail(MailFormModel formModel)
        {
            // Create the mail content as an object
            var mailContent = CreateMailContent(formModel);

            // Prepare the MonsiMail model
            SendingMonsiMail monsiModel = new SendingMonsiMail
            {
                receiverDid = formModel.DID,
                senderDid = "did:example:university", // Hardcoded sender DID
                mail = mailContent,          // Using the mail object, not stringified
            };

            // Serialize the MonsiMail to JSON
            string jsonPayload = JsonSerializer.Serialize(monsiModel, new JsonSerializerOptions
            {
                // DefaultIgnoreCondition = JsonIgnoreCondition.WhenWritingNull,
                WriteIndented = true // For readability, optional
            });

            // Replace encoded '+' characters
            jsonPayload = jsonPayload.Replace("\\u002B", "+");
            jsonPayload = jsonPayload.Replace("context", "@context");

            // Call API to encrypt the email content
            string encryptedContent;
            using (var httpClient = new HttpClient())
            {
                var requestContent = new StringContent(jsonPayload, Encoding.UTF8, "application/json");
                var response = await httpClient.PostAsync("http://localhost:80/api/mail/new", requestContent);

                if (response.IsSuccessStatusCode)
                {
                    encryptedContent = await response.Content.ReadAsStringAsync(); // Encrypted content is the response body
                }
                else
                {
                    ViewBag.Error = $"Error encrypting email content: {response.StatusCode}";
                    return View("Index");
                }
            }

            // Create the email message with the encrypted content as the body
            var message = new MimeMessage();
            message.From.Add(new MailboxAddress("Me Mario", "qwertz0014@web.de"));  // Sender info
            message.To.Add(new MailboxAddress("", formModel.RecipientEmail));        // Recipient info
            message.Subject = "Monsimail";                                          // Hardcoded subject

            // Use the encrypted content as the email text body
            message.Body = new TextPart("plain")
            {
                Text = encryptedContent // API response is the email body
            };

            // Send the email using SMTP
            using (var client = new SmtpClient())
            {
                try
                {
                    // Connect to the SMTP server
                    await client.ConnectAsync("smtp.web.de", 587, MailKit.Security.SecureSocketOptions.StartTls);
                    await client.AuthenticateAsync("qwertz0014@web.de", "SSITe5tM@il"); // SMTP credentials

                    // Send the email
                    await client.SendAsync(message);

                    // Disconnect from the SMTP server
                    await client.DisconnectAsync(true);

                    ViewBag.Message = "Email sent successfully!";
                }
                catch (Exception ex)
                {
                    ViewBag.Error = $"Error sending email: {ex.Message}";
                }
            }

            return View("Index");
        }

        private object CreateMailContent(MailFormModel formModel)
        {
            // Construct the mail content with a hardcoded VC
            return new
            {
                subject = formModel.Subject,
                content = formModel.Body,
                vcs = new[]
                {
            new
            {
                // Hardcoded VC structure
                context = new[] {
                    "https://www.w3.org/ns/credentials/v2",
                    "https://www.w3.org/ns/credentials/examples/v2"
                },
                id = "http://university.example/credentials/3732",
                type = new[] { "VerifiableCredential", "ExamPariticipationConfirmation" },
                issuer = "did:example:university",
                validFrom = "2010-01-01T19:23:24Z",
                validUntil = "2020-01-01T19:23:24Z",
                credentialSubject = new
                {
                    id = formModel.DID, // Using the receiver DID from the form
                    authorization = new { type = "", name = "" }
                },
                proof = new
                {
                    type = "DataIntegrityProof",
                    proofValue = "dUXMAC/5kRTyUm2Cx1uaP9MWBZGpAryhTEAP0rkRFkmoU806Y5YCKPkcb4jZtv87OfhNs+tp3EQyz2PCSVIrSFPtG70JYWRnJEmQPgfvhrwdYdcC/vPiW8H+SiUAc8Z+Gu8eKn6fjwnVzWvBWaMzBkEDj2m0Vna9KOCYLYLmFMEIpOCLmw67ywkLIqkiIABt2EPyZ38f81D8xu9IjTnxCcvhpXzmqBkRXUP8kjKGiBlcsY4LQ+dxsbLzoTMsc36JXERi46Yw0kVn3Doavvy6HC5ajo8yLdFVI1eaJrN4y6IOYafknBfq0E3mV2VbAHakG6sdYcPWEKq+SvNpH5VOIw=="
                }
            }
        }
            };
        }
    }
}
