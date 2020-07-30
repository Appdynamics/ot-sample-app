const express = require('express'),
  http = require('http'),
  pino = require('pino'),
  expressPino = require('express-pino-logger'),
  bodyParser = require('body-parser');

var appd = require("appdynamics");
appd.profile({
      controllerHostName: 'apmqe-docker-03.corp.appdynamics.com',
      controllerPort: 8080,
      controllerSslEnabled: false,
      accountName: 'customer1',
      accountAccessKey: 'SJ5b2m7d1$354',
      applicationName: 'bookings',
      tierName: 'booking',
      nodeName: '0',
      debug: true,
      opentelemetry: {
        url: process.env.OTC_HOST
      },
      libagent: true,
      'logging': {
              'logfiles': [
                    {
                          'filename': 's_%N.log',
                          'level': 'DEBUG'
                    },
                    {
                          'filename': 'proxy_%N.protolog',
                          'level': 'TRACE',
                          'channel': 'protobuf'
                    }
                    ]           
                }
});

const logger = pino({
  level: process.env.LOG_LEVEL || 'info'
});
const expressLogger = expressPino({ logger });
const PORT = process.env.PORT || 5000;
const app = express();

const RESERVATION_SVC_URL = process.env.RSV_SVC;
const PAYMENT_SVC_URL = process.env.PAY_SVC;

app.use(expressLogger);
app.use(bodyParser.json());

app.post('/booking', (req, res) => {
  // Make a sync call out to reservation service
  logger.debug("Request body is %s", JSON.stringify(req.body));
  let reservationPayload = JSON.stringify({
    'date': req.body.date,
    'name': req.body.name
  });
  makeHTTPCall(RESERVATION_SVC_URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Content-Length': reservationPayload.length
    }
  }, reservationPayload, (err, reservationResult) => {
    if (!err) {
      // Make a sync call out to payment service
      let paymentPayload = JSON.stringify({
        'card': req.body.card
      });
      makeHTTPCall(PAYMENT_SVC_URL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Content-Length': paymentPayload.length
        }
      }, paymentPayload, (err, paymentResult) => {
        if (!err) {
          res.statusCode = 200;
          res.json({ payment: paymentResult, reservation: reservationResult });
        } else {
          logger.error('Payment service call failed due to %s', err);
          res.statusCode = 400;
          res.send('Payment Call Failed');
        }
      });
    } else {
      logger.error('Reservation service call failed due to %s', err);
      res.statusCode = 400;
      res.send('Reservation Call Failed');
    }
  })
});

app.listen(PORT, () => {
  logger.info('Server is running on port %d', PORT);
});

function makeHTTPCall(url, options, payload, cb) {
  const req = http.request(url, options, (response) => {
    let body = '';
    response.on('data', (chunk) => {
      body += chunk;
    });
    response.on('end', () => {
      logger.debug('Response is %s', body);
      if (response.statusCode == 200) {
        cb(null, JSON.parse(body));
      } else {
        cb(body);
      }
    });
  }).on('error', (err) => {
    cb(err);
  });
  if (payload)
    req.write(payload);
  req.end();
}
