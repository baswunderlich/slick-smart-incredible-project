using System;
using System.Collections.Generic;
using System.Windows.Forms;

namespace Janus
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();

            // Anpassen und Initialisieren der ListView beim Start
            InitializeListView();
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
            // Füge hier deine E-Mail-Daten hinzu
            var emails = new List<(string Subject, string Sender, string Date, string Content)>
        {
            ("Betreff 1", "absender@example.com", "01.01.2024", "Inhalt der E-Mail 1"),
            ("Betreff 2", "absender2@example.com", "02.01.2024", "Inhalt der E-Mail 2"),
            ("Betreff 3", "absender2@example.com", "02.01.2024", "Inhalt der E-Mail 3"),
            ("Betreff 4", "absender2@example.com", "02.01.2024", "Inhalt der E-Mail 4"),
            ("Betreff 5", "absender2@example.com", "02.01.2024", "Inhalt der E-Mail 5"),
            ("Betreff 6", "absender2@example.com", "02.01.2024", "Inhalt der E-Mail 6"),
            ("Betreff 7", "absender2@example.com", "02.01.2024", "Inhalt der E-Mail 7"),
            ("Betreff 8", "absender2@example.com", "02.01.2024", "Inhalt der E-Mail 8"),
            ("Betreff 9", "absender2@example.com", "02.01.2024", "Inhalt der E-Mail 9"),
            ("Betreff 10", "absender2@example.com", "02.01.2024", "Inhalt der E-Mail 10"),
        };

            listViewEmails.Items.Clear();
            foreach (var email in emails)
            {
                var item = new ListViewItem(new[] { email.Subject, email.Sender, email.Date });
                item.Tag = email.Content;  // Speichert den E-Mail-Inhalt im Tag für späteren Zugriff
                listViewEmails.Items.Add(item);
            }
        }

        private void listViewEmails_SelectedIndexChanged(object sender, EventArgs e)
        {
            if (listViewEmails.SelectedItems.Count == 1)
            {
                var selectedItem = listViewEmails.SelectedItems[0];
                selectedMailContentBox.Text = (string)selectedItem.Tag;  // Zeige den Inhalt der ausgewählten E-Mail an
            }
        }
    }
}
