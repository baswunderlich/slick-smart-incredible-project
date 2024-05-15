using Janus.Models;
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
            // Verbinde Event-Handler
            listViewEmails.SelectedIndexChanged += listViewEmails_SelectedIndexChanged;

            // Lade E-Mail-Liste beim Start
            FillEmailList();
        }

        private void InitializeListView()
        {
            // Stelle sicher, dass die ListView leer ist
            listViewEmails.Columns.Clear();

            // Füge eine Spalte hinzu
            ColumnHeader columnHeader = new ColumnHeader();
            columnHeader.Text = "E-Mails";  // Text der Spalte
            columnHeader.Width = -2;        // Automatische Breite
            listViewEmails.Columns.Add(columnHeader);

            // Optional: Stelle den View-Modus auf 'Details' ein
            listViewEmails.View = View.Details;

            // Füge Beispieldaten hinzu (optional, für Testzwecke)
            FillEmailList();
        }

        private void FillEmailList()
        {
            // Test mail data
            var emails = new List<Email>
            {
                new Email { Subject = "Betreff 1", Sender = "absender@example.com", Date = "01.01.2024", Content = "Inhalt der E-Mail 1" },
                new Email { Subject = "Betreff 2", Sender = "absender@example.com", Date = "01.01.2024", Content = "Inhalt der E-Mail 2" },
                new Email { Subject = "Betreff 3", Sender = "absender@example.com", Date = "01.01.2024", Content = "Inhalt der E-Mail 3" },
                new Email { Subject = "Betreff 4", Sender = "absender@example.com", Date = "01.01.2024", Content = "Inhalt der E-Mail 4" },
                new Email { Subject = "Betreff 5", Sender = "absender@example.com", Date = "01.01.2024", Content = "Inhalt der E-Mail 5" },
                new Email { Subject = "Betreff 6", Sender = "absender@example.com", Date = "01.01.2024", Content = "Inhalt der E-Mail 6" },
            };

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
            e.DrawDefault = false;  // Wichtig, um die Standardzeichnung zu deaktivieren
        }

        private void ListViewEmails_DrawSubItem(object sender, DrawListViewSubItemEventArgs e)
        {
            if (e.ColumnIndex == 0)
            {
                // E-Mail-Daten holen
                var email = (Email)e.Item.Tag; // Stelle sicher, dass du das Tag entsprechend setzt

                // Erste Zeile: Absender und Datum
                string firstLine = $"{email.Sender}";
                string date = email.Date;
                string secondLine = email.Subject;

                // Textformatierung
                StringFormat sfLeft = new StringFormat();
                sfLeft.LineAlignment = StringAlignment.Near;
                sfLeft.Alignment = StringAlignment.Near;

                StringFormat sfRight = new StringFormat();
                sfRight.LineAlignment = StringAlignment.Near;
                sfRight.Alignment = StringAlignment.Far;

                // Zeichne erste Zeile: Absender und Datum
                e.Graphics.DrawString(firstLine, e.Item.Font, Brushes.Black, e.Bounds, sfLeft);
                e.Graphics.DrawString(date, e.Item.Font, Brushes.Black, e.Bounds, sfRight);

                // Zeichne zweite Zeile: Betreff
                Rectangle secondLineRect = new Rectangle(e.Bounds.Left, e.Bounds.Top + e.Item.Font.Height + 2, e.Bounds.Width, e.Item.Font.Height);
                e.Graphics.DrawString(secondLine, e.Item.Font, Brushes.Black, secondLineRect, sfLeft);
            }
        }
    }
}
