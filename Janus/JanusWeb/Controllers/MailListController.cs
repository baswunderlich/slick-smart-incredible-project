using JanusWeb.Services;
using JanusWeb.Models;
using Microsoft.AspNetCore.Mvc;
using System.Net.Http;

namespace JanusWeb.Controllers
{
    public class MailListController : Controller
    {
        public IActionResult Index()
        private readonly HttpClient _httpClient;

        public MailListController(HttpClient httpClient)
        {
            _httpClient = httpClient;
        }
        {
            MailList maillist = new MailList();
            EmailService emailService = new EmailService("imap.web.de", 993, true, "qwertz0014@web.de", "SSITe5tM@il");
            maillist.EmailList = emailService.FetchEmails();


            return View(maillist);
        }
    }
}
