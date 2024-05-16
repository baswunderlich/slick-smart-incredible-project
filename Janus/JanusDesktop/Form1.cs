using Janus.Models;
using Janus.Services;
using System;
using System.Collections.Generic;
using System.Drawing;
using System.Windows.Forms;
using static System.Windows.Forms.VisualStyles.VisualStyleElement.ListView;

namespace Janus
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();

            // Anpassen und Initialisieren der ListView beim Start
            InitializeListView();

            listViewEmails.OwnerDraw = true;
            listViewEmails.DrawColumnHeader += ListViewEmails_DrawColumnHeader;
            listViewEmails.DrawItem += ListViewEmails_DrawItem;
            listViewEmails.DrawSubItem += ListViewEmails_DrawSubItem;
            listViewEmails.SelectedIndexChanged += listViewEmails_SelectedIndexChanged;

            // Lade E-Mail-Liste beim Start
            FillEmailList();
        }

        private void InitializeListView()
        {
            // Setup the column
            listViewEmails.View = View.Details;
            listViewEmails.Columns.Add("E-Mails", -2);  // Breite automatisch anpassen
            listViewEmails.FullRowSelect = true;

            // Add imagelist for resizing rows
            ImageList imgList = new ImageList();
            imgList.ImageSize = new Size(1, 40);  // Breite ist irrelevant, Höhe setzt die Zeilenhöhe
            listViewEmails.SmallImageList = imgList;
        }

        private void FillEmailList()
        {
            //// Test mail data
            //var emails = new List<Email>
            //{
            //    new Email { Subject = "Betreff 1", Sender = "absender@example.com", Date = "01.01.2024", Content = "Inhalt der E-Mail 1" },
            //    new Email { Subject = "Betreff 2", Sender = "absender@example.com", Date = "01.01.2024", Content = "Inhalt der E-Mail 2" },
            //    new Email { Subject = "Betreff 3", Sender = "absender@example.com", Date = "01.01.2024", Content = "Inhalt der E-Mail 3" },
            //    new Email { Subject = "Betreff 4", Sender = "absender@example.com", Date = "01.01.2024", Content = "Inhalt der E-Mail 4" },
            //    new Email { Subject = "Betreff 5", Sender = "absender@example.com", Date = "01.01.2024", Content = "Inhalt der E-Mail 5" },
            //    new Email { Subject = "Betreff 6", Sender = "absender@example.com", Date = "01.01.2024", Content = "Inhalt der E-Mail 6" },
            //};
            EmailService emailService = new EmailService("imap.web.de", 993, true, "qwertz0014@web.de", "SSITe5tM@il");
            var emails = emailService.FetchEmails();

            listViewEmails.Items.Clear();
            foreach (var email in emails)
            {
                var item = new ListViewItem(new[] { email.Subject }); // Zeige nur den Betreff in der Liste an
                item.SubItems.Add(email.Sender);
                item.SubItems.Add(email.Date);
                item.Tag = email;  // Speichert das gesamte E-Mail-Objekt im Tag für späteren Zugriff
                listViewEmails.Items.Add(item);
            }
        }

        private void listViewEmails_SelectedIndexChanged(object sender, EventArgs e)
        {
            if (listViewEmails.SelectedItems.Count == 1)
            {
                ListViewItem selectedItem = listViewEmails.SelectedItems[0];
                Email email = (Email)selectedItem.Tag;  // Cast das Tag-Objekt zurück zu einem Email-Objekt
                selectedMailContentBox.Text = email.Content;  // Zeige den Inhalt der ausgewählten E-Mail an
            }
            else
            {
                selectedMailContentBox.Text = "";  // Leere den Inhalt, wenn kein Item ausgewählt ist
            }
        }

        private void ListViewEmails_DrawColumnHeader(object sender, DrawListViewColumnHeaderEventArgs e)
        {
            e.DrawDefault = true;
        }

        private void ListViewEmails_DrawItem(object sender, DrawListViewItemEventArgs e)
        {
            e.DrawDefault = false; // Wichtig, um die Standardzeichnung zu deaktivieren
        }

        private void ListViewEmails_DrawSubItem(object sender, DrawListViewSubItemEventArgs e)
        {
            Email email = (Email)e.Item.Tag;

            // Textformatierungen
            StringFormat sfLeft = new StringFormat { Alignment = StringAlignment.Near, LineAlignment = StringAlignment.Center };
            StringFormat sfRight = new StringFormat { Alignment = StringAlignment.Far, LineAlignment = StringAlignment.Center };

            // Berechne die Rechtecke für Texte
            int splitPoint = e.Bounds.Width / 2; // Teilt die Breite für Absender und Datum
            Rectangle leftRect = new Rectangle(e.Bounds.Left, e.Bounds.Top, splitPoint, e.Bounds.Height / 2);
            Rectangle rightRect = new Rectangle(e.Bounds.Left + splitPoint, e.Bounds.Top, splitPoint, e.Bounds.Height / 2);
            Rectangle secondLineRect = new Rectangle(e.Bounds.Left, e.Bounds.Top + e.Bounds.Height / 2, e.Bounds.Width, e.Bounds.Height / 2);

            // Zeichne den Text
            e.Graphics.DrawString(email.Sender, e.Item.Font, Brushes.Black, leftRect, sfLeft);
            e.Graphics.DrawString(email.Date, e.Item.Font, Brushes.Black, rightRect, sfRight);
            e.Graphics.DrawString(email.Subject, e.Item.Font, Brushes.Black, secondLineRect, sfLeft);
        }

        private void selectedMailContentBox_TextChanged(object sender, EventArgs e)
        {

        }
    }
}
