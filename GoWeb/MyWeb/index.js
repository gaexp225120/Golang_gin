const express = require('express')
const app = express()
const port = 3000

app.use('/', express.static('test'));
app.use('/test', express.static('test'));


app.listen(port, () => {
  console.log(`app listening at http://localhost:${port}`)
})