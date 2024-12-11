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
        private readonly HttpClient _httpClient;
        private readonly IConfiguration _configuration;

        public MailCreatorController(HttpClient httpClient, IConfiguration configuration)
        {
            _httpClient = httpClient;
            _configuration = configuration;
        }

        [HttpGet]
        public IActionResult Index()
        {
            return View();  // Return the page to create and send an email
        }

        [HttpPost]
        public async Task<IActionResult> FetchVCs([FromBody] FetchVCsRequest request)
        {
            var did = request.SenderDid;
            try
            {
                var response = await _httpClient.PostAsync("http://localhost:80/api/vc",
                    new StringContent($"{{\"did\":\"{did}\"}}", Encoding.UTF8, "application/json"));

                if (response.IsSuccessStatusCode)
                {
                    var vcsJson = await response.Content.ReadAsStringAsync();
                    return Content(vcsJson, "application/json");
                }
                else
                {
                    return StatusCode((int)response.StatusCode, $"Error fetching VCs: {response.ReasonPhrase}");
                }
            }
            catch (Exception ex)
            {
                return StatusCode(500, $"Internal server error: {ex.Message}");
            }
        }

        [HttpPost]
        public async Task<IActionResult> SendEmail(MailFormModel formModel)
        {
            // Create the mail content as an object
            var mailContent = CreateMailContent(formModel);

            // Prepare the MonsiMail model
            SendingMonsiMail monsiModel = new SendingMonsiMail
            {
                receiverDid = formModel.ReceiverDID,
                senderDid = formModel.SenderDID,
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
            jsonPayload = jsonPayload.Replace("\\u0022", "\"");
            jsonPayload = jsonPayload.Replace("\"[", "");
            jsonPayload = jsonPayload.Replace("]\"", "");
            //jsonPayload = jsonPayload.Replace("context", "@context");

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
            message.From.Add(new MailboxAddress("Me Mario", _configuration["EmailSettings:Email"]));  // Sender info
            message.To.Add(new MailboxAddress("", formModel.RecipientEmail));        // Recipient info
            message.Subject = "monsimail";                                          // Hardcoded subject

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
                    await client.ConnectAsync(_configuration["ServerSettings:SMTPServer"], Convert.ToInt32(_configuration["ServerSettings:SMTPPort"]), MailKit.Security.SecureSocketOptions.StartTls);
                    await client.AuthenticateAsync(_configuration["EmailSettings:Email"], _configuration["EmailSettings:Password"]); // SMTP credentials

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
            var selectedVCs = formModel.SelectedVCs;

            return new
            {
                subject = formModel.Subject,
                content = formModel.Body,
                vcs = selectedVCs // Use the selected VCs from the frontend
            };
        }
    }
}
