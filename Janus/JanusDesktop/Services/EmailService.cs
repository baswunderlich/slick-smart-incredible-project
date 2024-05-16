using JanusDesktop.Models;
using MailKit;
using MailKit.Net.Imap;
using System.Collections.Generic;
using System.Linq;


namespace JanusDesktop.Services
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
                    emails.Add(new Email
                    {
                        Subject = message.Subject,
                        Sender = message.From.Mailboxes.FirstOrDefault()?.Address,
                        Date = message.Date.DateTime.ToString(),
                        Content = message.TextBody
                    });
                }

                client.Disconnect(true);
            }

            return emails;
        }
    }
}
