import { server } from "../const/apiBase";
import axios from 'axios';

export const handleClick = () => {
  axios.get(server)
    .then((res) => {
      console.log(res);
      console.log(res.data.book);
      return res.data.book;
    })
    .catch(console.error);
}
