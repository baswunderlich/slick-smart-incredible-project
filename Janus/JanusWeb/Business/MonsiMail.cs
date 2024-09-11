namespace JanusWeb.Business
{
    public class MonsiMail
    {
        public string receiverDid { get; set; }
        public string senderDid { get; set; }
        public string? AESKey { get; set; }
        public string mail { get; set; }
        public string? signature { get; set; }
    }
    public class SendingMonsiMail
    {
        public string receiverDid { get; set; }
        public string senderDid { get; set; }
        public object mail { get; set; }
    }
}
