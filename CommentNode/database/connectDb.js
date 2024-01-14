
const mongoose = require('mongoose');

const conncetDb = mongoose.connect('mongodb://127.0.0.1:27017/commentdb', {
    useNewUrlParser: true,
    useUnifiedTopology: true,
});

module.exports = mongoose.mongo(conncetDb)
