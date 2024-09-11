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
                    if (jsonContent == null)
                    {
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

                    email.Did = mailContent.receiverDid;
                    email.Signature = mailContent.signature;


                    string requestJson = JsonSerializer.Serialize(mailContent);
                    StringContent httpContent = new StringContent(requestJson, Encoding.UTF8, "application/json");

                    HttpResponseMessage? response = null;
                    try
                    {
                        // Send the POST request to the API
                        response = await _httpClient.PostAsync("http://localhost:80/api/mail", httpContent);
                    }
                    catch (Exception)
                    {
                        ViewBag.Error = $"API request failed, is the wallet offline?";
                    }

                    if (response != null && response.IsSuccessStatusCode)
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
                            CheckVcsForValidityAndSetInMail(email, decryptResponse);
                            email.Content = $"VC check: {email.VCValid}\n\n{decryptResponse.content}\n\n ----original encrypted mail---- \n\n{email.Content}";
                            email.Subject += $": {decryptResponse.subject}";
                        }
                        else
                        {
                            ViewBag.Error = $"Failed to parse decryption response; original response: {responseContent}";
                        }
                    }
                    else if (response == null)
                    {
                        ViewBag.Error = $"no response!";
                    }
                    else if (!response.IsSuccessStatusCode)
                    {
                        ViewBag.Error = $"Response not successfull, status code: {response.StatusCode}";
                    }
                }

            }

            return View(maillist);
        }

        private static void CheckVcsForValidityAndSetInMail(Email email, DecryptResponse decryptResponse)
        {
            int validVCCounter = 0;
            string vcValid = "";

            foreach (var vc in decryptResponse.reviewedVCs)
            {
                if (vc.TryGetProperty("monsiValid", out JsonElement monsiValid))
                {
                    bool isValid = monsiValid.GetBoolean();

                    if (!isValid)
                    {
                        // Serialize the VC JsonElement back to a string to display it
                        vcValid += $"The VC is not valid! {vc.GetRawText()}<br/>";
                    }
                    else
                    {
                        validVCCounter++;
                    }
                }
            }

            if (validVCCounter == decryptResponse.reviewedVCs.Count)
            {
                vcValid = "All VCs are valid!";
            }
            email.VCValid = vcValid;
        }
    }
}