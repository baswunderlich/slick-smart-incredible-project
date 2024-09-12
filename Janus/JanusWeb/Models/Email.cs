﻿namespace JanusWeb.Models
{
    public class Email
    {
        public DateTime Date { get; set; }
        public string Sender { get; set; }
        public string Subject { get; set; }
        public string Content { get; set; }
        public string? Did { get; set; }
        public string? Signature { get; set; }
        public string VCValid { get; set; }
    }
}
