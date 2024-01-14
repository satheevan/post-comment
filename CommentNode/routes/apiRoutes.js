const express = require('express');
const router = express.Router();
const commentModel = require('../models/commentModel');
const commentController = require('../controller/createCommentcontroller');

//create a new comment
router.post('/createcomment', commentController.createComment);

//Get all /Retrive all
router.get('/getallcomment', (req, res) => {

    commentModel.find({}, (err, comments) => {
        if (err) {
            res.status(500).json({ error: err.message })
        } else {
            res.json(comments)
        }
    });
})

//get all comments of a specfic postId /Retrive all
router.get('/get-comments/:postId', async (req, res) => {
    console.log(req.params); 

    if (!req.params['postId']) {
        res.status(400).json({ error: 'PostId is invalid' });
        console.log(error);
    } else {
        try {
            const comments = await commentModel.find({ 'postId': req.params['postId'] })
            res.json(comments)
        } catch (err) {
            res.status(500).json({ error: err.message })
        }
    }
})

module.exports = router;