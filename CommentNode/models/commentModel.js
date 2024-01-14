const mongoose = require('mongoose');

const commentSchema = new mongoose.Schema({
    id: {
        type: Number,
        // Automation: true,
        // required:true,
    },
    description: String,
    postId: {
        type: Number,
        required: true,
    },
    userId: {
        type: Number,
        required: true
    },
    createAt: {
        type: Date,
        default: Date.now
    },
    modifiedAt: {
        type: Date,
    },
    deletedAt: {
        type: Date,
    },

})

module.exports = mongoose.model('comment', commentSchema)