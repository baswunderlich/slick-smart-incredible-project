using JanusWeb.Services;
using JanusWeb.Models;
using Microsoft.AspNetCore.Mvc;

namespace JanusWeb.Controllers
{
    public class MailListController : Controller
    {
        public IActionResult Index()
        {
            MailList maillist = new MailList();
            EmailService emailService = new EmailService("imap.web.de", 993, true, "qwertz0014@web.de", "SSITe5tM@il");
            maillist.EmailList = emailService.FetchEmails();


            return View(maillist);
        }
    }
}
