import { useForm } from "@mantine/form";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import {
  createTask,
  getTaskList,
} from "../gen/rpc/task/v1/task-TaskService_connectquery";

type FormValue = {
  name: string;
};

export const useNewTask = () => {
  const client = useQueryClient();
  const createMutation = useMutation(createTask.useMutation());
  const form = useForm<FormValue>({
    initialValues: {
      name: "",
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    createMutation
      .mutateAsync({
        name: values.name,
      }, {
        onSuccess: () => {
          client
            .refetchQueries(getTaskList.useQuery())
            .catch((e) => console.error(e));
        },
      })
      .catch((e) => console.error(e));
    form.reset();
  });

  return { form, createMutation, handleSubmit}
}
