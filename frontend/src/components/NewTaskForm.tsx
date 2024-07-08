import { Alert, Box, Button, Group, TextInput } from "@mantine/core";
import { FC } from "react";
import { useNewTask } from "../hooks/UseNewTask";

export const NewTaskForm: FC = () => {
  const { form, createMutation, handleSubmit } = useNewTask();

  return (
    <Box>
      <form onSubmit={handleSubmit}>
        {createMutation.isError && (
          <Alert title="Error" color="red">
            {createMutation.error.message}
          </Alert>
        )}
        <Group mt="md">
          <TextInput
            {...form.getInputProps("name")}
            label="New Task"
            autoComplete="off"
            placeholder="New Task"
          />
          <Button type="submit" loading={createMutation.isPending}>
            Add
          </Button>
        </Group>
      </form>
    </Box>
  )
}
