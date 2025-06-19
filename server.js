// JavaScript source code
const express = require("express");
const bodyParser = require("body-parser");
const cors = require("cors");
const { Generate } = require("./pkg/imageGO/imageGO"); // caminho para a função

const app = express();
app.use(cors());
app.use(bodyParser.json());

app.post("/generate", async (req, res) => {
    try {
        const result = await Generate(req.body);
        res.json(result);
    } catch (err) {
        console.error(err);
        res.status(500).json({ error: "Image generation failed" });
    }
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`✅ Server running on port ${PORT}`);
});
