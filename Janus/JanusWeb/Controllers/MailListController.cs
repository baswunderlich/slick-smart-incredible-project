using JanusWeb.Services;
using JanusWeb.Models;
using Microsoft.AspNetCore.Mvc;
using System.Text.Json;
using JanusWeb.Business;
using System.Net.Http;
using System.Text;

namespace JanusWeb.Controllers
{
    public class MailListController : Controller
    {
        private readonly HttpClient _httpClient;

        public MailListController(HttpClient httpClient)
        {
            _httpClient = httpClient;
        }
        public async Task<IActionResult> Index()
        {
            MailList maillist = new MailList();
            EmailService emailService = new EmailService("imap.web.de", 993, true, "qwertz0014@web.de", "SSITe5tM@il");
            maillist.EmailList = emailService.FetchEmails();

            maillist.EmailList = maillist.EmailList.OrderByDescending(email => email.Date).ToList();

            // extract the did from the mail body, empty string if not found
            foreach (var email in maillist.EmailList)
            {
                if (email.Subject.Contains("monsi"))
                {
                    // Example JSON content stored as a string
                    string jsonContent = email.Content;
                    if (jsonContent == null) {
                        ViewBag.Error = $"no content in mail";
                        continue;
                    }

                    // Deserialize the JSON string to an object
                    MonsiMail? mailContent = JsonSerializer.Deserialize<MonsiMail>(jsonContent);

                    if (mailContent == null) // parsing failed
                    {
                        ViewBag.Error = $"parsing failed; original mail content: {jsonContent}";
                        continue; 
                    }

                    email.Did = mailContent.did;
                    email.Signature = mailContent.signature;

                    // Create a DecryptRequest object
                    DecryptRequest decryptRequest = new DecryptRequest
                    {
                        did = mailContent.did,
                        content = mailContent.orgMail
                    };

                    // Convert the request object to JSON
                    string requestJson = JsonSerializer.Serialize(decryptRequest);
                    StringContent httpContent = new StringContent(requestJson, Encoding.UTF8, "application/json");

                    // Send the POST request to the API
                    HttpResponseMessage response = await _httpClient.PostAsync("http://localhost:80/api/decrypt", httpContent);

                    if (response.IsSuccessStatusCode)
                    {

                        // Read and deserialize the response content
                        string responseContent = await response.Content.ReadAsStringAsync();
                        // Check if the response is a stringified JSON (double-encoded)
                        if (responseContent.StartsWith("\"") && responseContent.EndsWith("\""))
                        {
                            // Remove quotes and unescape the string
                            responseContent = JsonSerializer.Deserialize<string>(responseContent);
                        }
                        DecryptResponse? decryptResponse = JsonSerializer.Deserialize<DecryptResponse>(responseContent);

                        if (decryptResponse != null)
                        {
                            email.Content = $"{decryptResponse.content}\n\n ----original encrypted mail---- \n\n{email.Content}";
                        }
                        else
                        {
                            ViewBag.Error = $"Failed to parse decryption response; original response: {responseContent}";
                        }
                    }
                    else
                    {
                        ViewBag.Error = $"API call failed with status code: {response.StatusCode}";
                    }
                }

            }

            return View(maillist);
        }
    }
}
