using HtmlAgilityPack;
using JanusWeb.Models;
using MailKit;
using MailKit.Net.Imap;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Text.RegularExpressions;


namespace JanusWeb.Services
{
    public class EmailService
    {
        private readonly string host;
        private readonly int port;
        private readonly bool useSsl;
        private readonly string username;
        private readonly string password;

        public EmailService(string host, int port, bool useSsl, string username, string password)
        {
            this.host = host;
            this.port = port;
            this.useSsl = useSsl;
            this.username = username;
            this.password = password;
        }

        public List<Email> FetchEmails()
        {
            var emails = new List<Email>();

            using (var client = new ImapClient())
            {
                client.Connect(host, port, useSsl);
                client.Authenticate(username, password);

                client.Inbox.Open(FolderAccess.ReadOnly);
                var results = client.Inbox.Fetch(0, -1, MessageSummaryItems.Full | MessageSummaryItems.UniqueId);

                foreach (var summary in results)
                {
                    var message = client.Inbox.GetMessage(summary.UniqueId);
                    string content = message.TextBody;

                    if (string.IsNullOrEmpty(content) && !string.IsNullOrEmpty(message.HtmlBody))
                    {
                        // Extract plain text from HTML
                        content = ConvertHtmlToPlainText(message.HtmlBody);
                    }

                    // Optional: Further sanitize the content if needed
                    content = SanitizeJsonString(content);

                    emails.Add(new Email
                    {
                        Subject = message.Subject,
                        Sender = message.From.Mailboxes.FirstOrDefault()?.Address!,
                        Date = message.Date.DateTime,
                        Content = content
                    });
                }

                client.Disconnect(true);
            }

            return emails;
        }

        private string ConvertHtmlToPlainText(string html)
        {
            var doc = new HtmlDocument();
            doc.LoadHtml(html);

            // Extract the inner text from the HTML and then decode the HTML entities
            string plainText = doc.DocumentNode.InnerText;

            // Decode HTML entities like &quot; to "
            plainText = WebUtility.HtmlDecode(plainText);

            // Replace non-breaking spaces with regular spaces
            plainText = plainText.Replace('\u00A0', ' ');

            // Optionally, trim the text to remove leading/trailing whitespace
            plainText = plainText.Trim();

            return plainText;
        }
        private string SanitizeJsonString(string json)
        {
            // Remove any extraneous characters or whitespace if necessary
            // For example, ensure it starts with { and ends with }
            json = json.Trim();

            if (!json.StartsWith("{"))
            {
                int start = json.IndexOf('{');
                if (start != -1)
                {
                    json = json.Substring(start);
                }
            }

            if (!json.EndsWith("}"))
            {
                int end = json.LastIndexOf('}');
                if (end != -1)
                {
                    json = json.Substring(0, end + 1);
                }
            }

            return json;
        }
    }
}
