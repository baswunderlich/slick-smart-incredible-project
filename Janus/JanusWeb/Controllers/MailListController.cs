using JanusWeb.Services;
using JanusWeb.Models;
using Microsoft.AspNetCore.Mvc;
using System.Text.Json;
using JanusWeb.Business;
using System.Text;

namespace JanusWeb.Controllers
{
    public class MailListController : Controller
    {
        private readonly HttpClient _httpClient;
        private readonly EmailService _emailService;

        public MailListController(HttpClient httpClient)
        {
            _httpClient = httpClient;
            _emailService = new EmailService("imap.web.de", 993, true, "qwertz0014@web.de", "SSITe5tM@il");
        }
        public async Task<IActionResult> Index()
        {
            MailList maillist = new MailList
            {
                EmailList = _emailService.FetchEmails()
            };

            maillist.EmailList = maillist.EmailList.OrderByDescending(email => email.Date).ToList();

            // extract the did from the mail body, empty string if not found
            foreach (var email in maillist.EmailList)
            {
                if (email.Subject.ToLower().Contains("monsi"))
                {
                    await ProcessMonsiMailAsync(email);
                }

            }

            return View(maillist);
        }

        private async Task ProcessMonsiMailAsync(Email email)
        {
            MonsiMail? mailContent = DeserializeEmailContent(email.Content);
            if (mailContent == null)
            {
                ViewBag.Error += "Failed to parse email content.";
                return;
            }

            email.Did = mailContent.receiverDid;
            email.Signature = mailContent.signature;


            var response = await PostToApiAsync(mailContent);

            if (response != null && response.IsSuccessStatusCode)
            {
                await HandleApiResponseAsync(response, email);
            }
            else
            {
                ViewBag.Error = $"API call failed with status code: {response?.StatusCode}";
            }
        }

        private MonsiMail? DeserializeEmailContent(string? content)
        {
            if (string.IsNullOrEmpty(content))
            {
                ViewBag.Error = "No content in email.";
                return null;
            }

            try
            {
                return JsonSerializer.Deserialize<MonsiMail>(content);
            }
            catch (Exception)
            {
                ViewBag.Error = $"Parsing failed; original mail content: {content}";
                return null;
            }
        }

        private async Task<HttpResponseMessage?> PostToApiAsync(MonsiMail mailContent)
        {
            var requestJson = JsonSerializer.Serialize(mailContent);
            var httpContent = new StringContent(requestJson, Encoding.UTF8, "application/json");

            try
            {
                return await _httpClient.PostAsync("http://localhost:80/api/mail", httpContent);
            }
            catch (Exception)
            {
                ViewBag.Error = "API request failed, is the wallet offline?";
                return null;
            }
        }

        private async Task HandleApiResponseAsync(HttpResponseMessage response, Email email)
        {
            var responseContent = await response.Content.ReadAsStringAsync();

            // Handle double-encoded JSON
            if (responseContent.StartsWith("\"") && responseContent.EndsWith("\""))
            {
                responseContent = JsonSerializer.Deserialize<string>(responseContent);
            }

            var decryptResponse = JsonSerializer.Deserialize<DecryptResponse>(responseContent);
            if (decryptResponse != null)
            {
                ProcessDecryptionResponse(email, decryptResponse);
            }
            else
            {
                ViewBag.Error = $"Failed to parse decryption response; original response: {responseContent}";
            }
        }

        private void ProcessDecryptionResponse(Email email, DecryptResponse decryptResponse)
        {
            string vcData = CheckVcsForValidityAndGetData(email, decryptResponse);
            email.OriginalMail = $"{email.Content}";
            email.Content = $"{decryptResponse.content}";
            email.Vcs = $"----VC check----\n\n {vcData}";
            email.Subject += $"{decryptResponse.subject}";
        }

        private static string CheckVcsForValidityAndGetData(Email email, DecryptResponse decryptResponse)
        {
            int validVCCounter = 0;
            string vcData = "";  // Contains the relevant data to each vc for the mail

            foreach (var vc in decryptResponse.reviewedVCs)
            {
                if (vc.TryGetProperty("monsiValid", out JsonElement monsiValid))
                {
                    bool isValid = monsiValid.GetBoolean();
                    string vcBasicData = GetVcBasicData(vc);  // Contains basic data like issuer, and validity dates

                    string vcNameObject = "";
                    vcNameObject += GetVcName(vc, "exam");
                    vcNameObject += GetVcName(vc, "authorization");
                    if (!isValid)
                    {
                        vcData += $"The VC is not valid!\n{vcNameObject}\n";
                    }
                    else
                    {
                        vcData += $"This vc is valid:\n{vcNameObject}\n";
                        validVCCounter++;
                    }
                    vcData += vcBasicData + "\n\n";
                }
            }
            string allvalid = "";
            if (validVCCounter == decryptResponse.reviewedVCs.Count)
            {
                allvalid = "All VCs are valid!\n\n";
            }
            return $"{allvalid}{vcData}";
        }
        }

        private static string GetVcName(JsonElement vc, string vcProperty)
        {
            if (vc.TryGetProperty("credentialSubject", out JsonElement credential)
                && credential.ValueKind == JsonValueKind.Object)
            {
                return GetVcPropertyString(credential, vcProperty) ?? "";
            }
            return "";
        }

        private static string GetVcBasicData(JsonElement vc)
        {
            string vcData = "";
            vcData += $"{GetVcPropertyString(vc, "issuer") ?? "unknown issuer"}\n";
            vcData += $"{GetVcPropertyString(vc, "validFrom") ?? "No Valid From date"}\n";
            vcData += $"{GetVcPropertyString(vc, "validUntil") ?? "No Valid Until date"}\n";
            return vcData;
        }

        private static string? GetVcPropertyString(JsonElement vc, string property)
        {
            if (vc.TryGetProperty(property, out JsonElement VCdata)
                && VCdata.ValueKind != JsonValueKind.Undefined)
            {
                return VCdata.ToString();
            }
            return null;
        }
    }
}