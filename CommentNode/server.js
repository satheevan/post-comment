const express = require('express')
const bodyParser = require('body-parser')

const mongoose = require('mongoose')

//folders Access
// const conncetDb = require('./database/connectDb')
const apiRouter = require('./routes/apiRoutes')

const app = express();

//MiddleWare
app.use(bodyParser.json());

//Database Connection
// manua
mongoose.connect("mongodb://127.0.0.1:27017/commentdb", {
    useNewUrlParser: true,
    useUnifiedTopology: true,
});
//export
// conncetDb()
const cors = require('cors');
app.use(cors({
    "origin": "*",
    "methods": "GET,HEAD,PUT,PATCH,POST,DELETE",
    "preflightContinue": false,
    "optionsSuccessStatus": 204
}))


//API Routes
app.use('/api', apiRouter);

const port = 3004;
app.listen(port, () => {
    console.log(`server is successfully running ${port}`);
});

