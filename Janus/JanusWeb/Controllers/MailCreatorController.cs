using MailKit.Net.Smtp;
using MimeKit;
using Microsoft.AspNetCore.Mvc;
using JanusWeb.Models;
using System.Text.Json;
using System.Text;

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
            // Verschlüsselten Inhalt vorbereiten
            var encryptionRequest = new
            {
                did = formModel.DID,
                content = $"{{\"subject\":\"{formModel.Subject}\",\"content\":\"{formModel.Body}\",\"vcs\":[{{\"monsiValid\":true}},{{\"monsiValid\":false}}]}}"
            };

            string encryptedContent = string.Empty;

            // API-Aufruf zur Verschlüsselung
            using (var httpClient = new HttpClient())
            {
                var requestContent = new StringContent(JsonSerializer.Serialize(encryptionRequest), Encoding.UTF8, "application/json");
                var response = await httpClient.PostAsync("http://localhost:80/api/encrypt", requestContent);

                if (response.IsSuccessStatusCode)
                {
                    string responseContent = await response.Content.ReadAsStringAsync();
                    encryptedContent = responseContent; // Encrypted response
                }
                else
                {
                    ViewBag.Error = $"Error encrypting email content: {response.StatusCode}";
                    return View("Index");
                }
            }

            // Erstelle die E-Mail mit verschlüsseltem Inhalt
            var message = new MimeMessage();
            message.From.Add(new MailboxAddress("me mario", "qwertz0014@web.de"));  // Change your email
            message.To.Add(new MailboxAddress("", formModel.RecipientEmail));
            message.Subject = formModel.Subject;

            // Verwende den verschlüsselten Inhalt als E-Mail-Text
            message.Body = new TextPart("plain")
            {
                Text = encryptedContent
            };

            // Senden der E-Mail
            using (var client = new SmtpClient())
            {
                try
                {
                    // Verbindung zum SMTP-Server herstellen
                    await client.ConnectAsync("smtp.web.de", 587, MailKit.Security.SecureSocketOptions.StartTls);
                    await client.AuthenticateAsync("qwertz0014@web.de", "SSITe5tM@il"); // Change these

                    // E-Mail senden
                    await client.SendAsync(message);

                    // Vom Server trennen
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
    }
}
