using System;

namespace PaymentService
{
    public class PaymentStatus
    {
        public DateTime Date { get; set; }

        public string Status { get; set; }

        public string TransactionId { get; set; }
    }
}
