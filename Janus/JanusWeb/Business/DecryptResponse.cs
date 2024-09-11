using System.Text.Json;

namespace JanusWeb.Business
{
    public class DecryptResponse
    {
        public bool signatureIsValid { get; set; }
        public string subject { get; set; }
        public string content { get; set; }
        public List<JsonElement> reviewedVCs { get; set; }
    }
}
