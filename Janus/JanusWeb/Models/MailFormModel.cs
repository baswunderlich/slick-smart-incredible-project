namespace JanusWeb.Models
{
    public class MailFormModel
    {
        public string RecipientEmail { get; set; }
        public string Subject { get; set; }
        public string SenderDID { get; set; }
        public string ReceiverDID { get; set; }
        public string Body { get; set; }
    }
}
