
const commentModel = require('../models/commentModel');

exports.createComment = async (req, res) => {
    //define the model 
    const createComment = new commentModel(req.body);
    // createComment.createAt = Date.now()

    //save 
    try {
        const result = await createComment.save();
        res.status(201).json(result)
    } catch (err) {
        res.status(400).json({ error: err.message });
    }
}