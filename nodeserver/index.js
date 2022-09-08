import express, { json } from 'express';
import cors from 'cors'

const PORT = 3000;
const app = express();

// app.use((req, res, next) => {
//   res.header('Access-Control-Allow-Origin', 'localhost:3000');
//   res.header(
//     'Access-Control-Allow-Headers',
//     'Origin, X-Requested-With, Content-Type, Accept'
//     );
//   next();
// });

app.use(cors());
  
app.use(json());

app.use(express.static('static'));

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});