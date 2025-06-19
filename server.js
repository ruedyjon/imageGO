const express = require("express");
const cors = require("cors");
const { execFile } = require("child_process");
const path = require("path");

const app = express();
app.use(cors());
app.use(express.json());

app.post("/generate", (req, res) => {
  const { prompt, authCode, modelName, aspectRatio, seed } = req.body;

  const args = [
    "-prompt", prompt,
    "-auth", authCode,
    "-model", modelName || "IMAGEN_3",
    "-aratio", aspectRatio || "IMAGE_ASPECT_RATIO_LANDSCAPE",
    "-seed", seed?.toString() || "42",
    "-name", "image" // Output will be image.png
  ];

  execFile(path.join(__dirname, "imagego"), args, (error, stdout, stderr) => {
    if (error) {
      console.error("Generation error:", stderr || error.message);
      return res.status(500).json({ error: "Generation failed" });
    }

    res.sendFile(path.join(__dirname, "image.png"));
  });
});

const port = process.env.PORT || 3000;
app.listen(port, () => {
  console.log(`ImageGO API running on port ${port}`);
});
