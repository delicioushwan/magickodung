import { useEffect, useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { ButtonGroup, Button } from "@mui/material";

import Axios from "../../utils/axios";

export default function AddQuestions() {
  const [categories, setCategory] = useState([]);
  useEffect(async () => {
    const res = await Axios.get("/category");
    setCategory(res.data);
  }, []);
  console.log(categories);
  const { control, handleSubmit, watch, setValue } = useForm({
    category: null,
  });
  const onSubmit = (data) => console.log(data);
  const { category } = watch();
  return (
    <div>
      <form onSubmit={handleSubmit(onSubmit)}>
        <ButtonGroup>
          {categories?.map((c) => (
            <Controller
              name={c.value}
              control={control}
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
        </ButtonGroup>
      </form>
      <p>로그인</p>
    </div>
  );
}
