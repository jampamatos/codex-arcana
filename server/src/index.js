const express = require('express');
const app = express();
const port = process.env.PORT || 3000;

app.get('/health', (_req, res) => {
    res.json({ status: 'ok'});
});

app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});