import { server } from "../const/apiBase";
import axios from 'axios';

export const handleClick = () => {
  axios.get(server)
    .then((res) => {
      console.log(res);
    })
    .catch(console.error);
}
