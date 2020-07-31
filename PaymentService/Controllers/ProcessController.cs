using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace PaymentService.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class ProcessController : ControllerBase
    {
        private readonly ILogger<ProcessController> _logger;

        public ProcessController(ILogger<ProcessController> logger)
        {
            _logger = logger;
        }

        [HttpPost]
        public PaymentStatus Post([FromBody] PaymentData paymentData)
        {
            var rng = new Random();
            return new PaymentStatus
            {
                Date = DateTime.Now,
                Status = "Charged " + paymentData.Card,
                TransactionId = "#" + rng.Next(100000, 999999)
            };
        }
    }
}
