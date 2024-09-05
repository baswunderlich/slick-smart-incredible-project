using System.Text.Json;

namespace JanusWeb.Business
{
    public class DecryptResponse
    {
        public string SUBJECT { get; set; }
        public string content { get; set; }
        public JsonElement vcs { get; set; } 
    }
}
