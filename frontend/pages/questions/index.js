//
// Copyright 2021 lemonade
//
import React, { useEffect, useState } from 'react';
import styled from 'styled-components';
import { Input, Button } from '@mui/material';
import { useForm } from 'react-hook-form';
import { useRouter } from 'next/router';

import Axios from '../../utils/axios';

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
  const router = useRouter();

  const nextQuestion = async () => {
    if (!window.confirm('다음?')) return;
    if (questions?.length > 1) {
      const newQuestions = questions.slice(1);
      setQuestions(newQuestions);
    } else {
      const res = await getQuestions();
      setQuestions(res);
    }
    setValue('answer', null);
  };

  const onSubmitAnswer = async () => {
    console.log({ answer });
    if (window.confirm('이것을 고르시겠습니까?')) {
      try {
        await postAnswer({ answer, questionId });
        window.alert('등록 성공~');
        await nextQuestion();
      } catch (e) {
        console.log({ e });
      }
    }
  };
  useEffect(async () => {
    console.log('ererere', window.location.search);
    const queryString = window.location.search;
    let params = new URLSearchParams(queryString);
    let q = parseInt(params.get('q'));
    if (Number.isNaN(q)) {
      const res = await getQuestions();
      setQuestions(res);
    } else {
      const res = await getQuestion({ q });
      setQuestions(res);
    }
  }, []);

  useEffect(() => {
    if (!questions?.[0]) return;
    setQuestion(questions[0]);
    console.log('runhere');
    console.log('questions[0].questionId', questions[0], questions[0]?.questionId);
    router.replace({ query: `q=${questions[0]?.questionId}` });
  }, [questions]);

  return (
    <div>
      <form onSubmit={handleSubmit(onSubmitAnswer)}>
        <div>질문 제목 : {title}</div>
        {options?.map(({ optionId, option, quantity }) => (
          <div key={optionId}>
            <Button
              type="button"
              variant={answer === option && 'contained'}
              onClick={() => setValue('answer', option)}
              color={answer === option ? 'secondary' : 'primary'}
              {...register('answer', { required: true })}
            >
              <div>
                <p>{option}</p>
                <img src="https://picsum.photos/200/300" />
              </div>
            </Button>
          </div>
        ))}
        <div>{errors?.answer?.type && '대답을 골라라~'}</div>
        <NextButton>
          <Button type="button" onClick={nextQuestion}>
            {`>`}
          </Button>
        </NextButton>
        <SubmitWrapper>
          <Button variant="contained" color="secondary" type="submit">
            답변하기
          </Button>
        </SubmitWrapper>
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
  const [comments, setComments] = useState([]);
  const onSubmitComment = () => {
    const { comment } = watch();
    console.log({ comment });
    if (window.confirm('댓글을 등록하시겠습니까?')) {
      window.alert('등록 성공~');
    }
  };

  useEffect(() => {
    console.log('comment init');
  }, []);

  return (
    <CommentWrapper>
      <div>댓글</div>
      <form onSubmit={handleSubmit(onSubmitComment)}>
        <Input {...register('comment', { required: true })} />
        <div>{errors?.comment?.type && '댓글을 입력하시오'}</div>
        <Button type="submit">댓글 남기기</Button>
      </form>
      <div>{comments.map((comment) => comment)}</div>
    </CommentWrapper>
  );
}

async function getQuestions() {
  const res = await Axios.get('/questions/');
  return res.data;
}

async function getQuestion({ q }) {
  const res = await Axios.get(`/questions/${q}`);
  return res.data;
}

async function postAnswer({ questionId, answer }) {
  const res = await Axios.post('/answers/', { questionId, answer });
  return res.data;
}

const NextButton = styled.div`
  position: fixed;
  top: 50%;
  right: 10px;
  & > button {
    font-size: 40px;
  }
`;

const SubmitWrapper = styled.div`
  padding: 20px;
  > button {
    width: 100%;
  }
`;

const CommentWrapper = styled.div`
  margin-top: 30px;
`;
