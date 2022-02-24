import { useEffect, useState } from "react";
import { useForm, Controller, useFieldArray } from "react-hook-form";
import { ButtonGroup, Button, Input } from "@mui/material";

import Axios from "../../utils/axios";

const initialValue = { category: null, option: [], title: "" };

export default function AddQuestions() {
  const [categories, setCategory] = useState([]);
  const {
    control,
    handleSubmit,
    watch,
    setValue,
    setError,
    reset,
    formState: { errors },
  } = useForm({ ...initialValue });
  const { category } = watch();
  const { fields, append, remove } = useFieldArray({
    control,
    name: "option",
  });
  console.log(errors);
  useEffect(async () => {
    const res = await Axios.get("/category");
    setCategory(res.data);
  }, []);

  const onSubmit = async (data) => {
    const { option, title, category } = data;
    const options = option.map(({ value }) => value);
    if (options.length < 2) {
      return setError("option", { type: "required" });
    }
    if (window.confirm("질문 ㄱ?")) {
      try {
        const res = await Axios.post("/questions/", { title, category, options });
        console.log(res);
        window.alert("성공");
        reset({ ...initialValue });
      } catch (err) {
        console.log(err);
      }
    }
  };
  console.log(watch());
  return (
    <div>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div>질문 제목!</div>
        <Controller
          name="title"
          control={control}
          rules={{ required: true }}
          render={({ field }) => <Input {...field} color={errors.title?.type ? "error" : ""} />}
        />
        {errors.title?.type === "required" && "제목 필수필수!!"}
        <ButtonGroup>
          {categories?.map((c) => (
            <Controller
              name="category"
              key={c.key}
              control={control}
              rules={{ required: true }}
              render={({ field }) => (
                <Button
                  key={c.key}
                  color={category === c.key ? "secondary" : "primary"}
                  onClick={() => setValue("category", c.key)}
                  {...field}
                >
                  {c.value}
                </Button>
              )}
            />
          ))}
          {errors.category?.type === "required" && "카테고리 선택 필수필수!!"}
        </ButtonGroup>
        <div>선택지!!!</div>
        {fields.map((item, index) => {
          return (
            <div key={item.id}>
              <Controller
                render={({ field }) => <Input variant="filled" {...field} />}
                name={`option.${index}.value`}
                rules={{ required: true }}
                control={control}
              />
              <Button type="button" color="secondary" onClick={() => remove(index)}>
                Delete
              </Button>
            </div>
          );
        })}
        {errors.option?.type === "required" && "선택지 최소 2개 이상!! 필수필수!!"}

        <section>
          <Button
            type="button"
            color="info"
            onClick={() => {
              append({ value: "" });
            }}
          >
            +
          </Button>
        </section>
        <Button type="submit">질문하기!</Button>
      </form>
    </div>
  );
}
