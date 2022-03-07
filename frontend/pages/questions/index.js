//
// Copyright 2021 lemonade
//
import React, { useEffect, useState } from "react";
import Axios from "../../utils/axios";
import { useForm } from "react-hook-form";

import { Input, Button } from "@mui/material";

export default function Questions({ Component, pageProps }) {
  const [questions, setQuestions] = useState([]);
  const [question, setQuestion] = useState(null);
  const {
    register,
    watch,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const onSubmitComment = () => {
    const { comment } = watch();
    console.log({ comment });
    if (window.confirm("댓글을 등록하시겠습니까?")) {
      window.alert("등록 성공~");
    }
  };

  const { questionId, options, title } = question || {};
  useEffect(async () => {
    const res = await getQuestions();
    setQuestions(res);
  }, []);

  useEffect(() => {
    setQuestion(questions[0]);
  }, [questions]);

  return (
    <div>
      <div key={questionId}>
        <div>{title}</div>
        {options?.map(({ optionId, option, quantity }) => (
          <div key={optionId}>{option}</div>
        ))}
      </div>

      <div>댓글</div>
      <form onSubmit={handleSubmit(onSubmitComment)}>
        <Input {...register("comment", { required: true })} />
        <div>{errors?.comment?.type && "댓글을 입력하시오"}</div>
        <Button type="submit">댓글</Button>
      </form>
    </div>
  );
}

async function getQuestions() {
  const res = await Axios.get("/questions/");
  return res.data;
}
