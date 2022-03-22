import { useEffect, useState } from 'react';
import Axios from '../../utils/axios';

export default function Login() {
  const [questions, setQuestions] = useState([]);

  useEffect(async () => {
    const res = await getQuestions();
    setQuestions(res);
  }, []);
  return (
    <div>
      <p>내 질문들</p>
      {questions?.map((question) => (
        <div>f</div>
      ))}
    </div>
  );
}

async function getQuestions() {
  const res = await Axios.get('/questions/user', { limit: 10, offset: 0 });
  return res.data;
}
