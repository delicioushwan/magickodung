//
// Copyright 2021 lemonade
//
import React, { useEffect, useState } from "react";
import Axios from "../../utils/axios";
import { useForm } from "react-hook-form";

import { Input, Button } from "@mui/material";

export default function Questions() {
  const [questions, setQuestions] = useState([]);
  const [question, setQuestion] = useState(null);
  const {
    register,
    watch,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm();

  const { answer } = watch();
  const { questionId, options, title } = question || {};

  const onSubmitAnswer = async () => {
    console.log({ answer });
    if (window.confirm("이것을 고르시겠습니까?")) {
      try {
        await postAnswer({ answer, questionId });
        window.alert("등록 성공~");
      } catch (e) {
        console.log(e);
      }
    }
  };
  useEffect(async () => {
    const res = await getQuestions();
    setQuestions(res);
  }, []);

  useEffect(() => {
    setQuestion(questions[0]);
  }, [questions]);

  return (
    <div>
      <form onSubmit={handleSubmit(onSubmitAnswer)}>
        <div>질문 제목 : {title}</div>
        {options?.map(({ optionId, option, quantity }) => (
          <div key={optionId}>
            <Button
              type="button"
              variant={answer === option && "contained"}
              onClick={() => setValue("answer", option)}
              color={answer === option ? "secondary" : "primary"}
              {...register("answer", { required: true })}
            >
              {option}
            </Button>
          </div>
        ))}
        <div>{errors?.answer?.type && "대답을 골라라~"}</div>
        <Button type="submit">답변하기</Button>
      </form>

      <Comments />
    </div>
  );
}

function Comments() {
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

  console.log(errors);
  return (
    <div>
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

async function postAnswer({ questionId, answer }) {
  const res = await Axios.post("/answers/", { questionId, answer });
  return res.data;
}
