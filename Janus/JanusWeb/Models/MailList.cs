namespace JanusWeb.Models
{
    public class MailList
    {
        public IEnumerable<Email> EmailList { get; set; } = new List<Email>();
    }
}
