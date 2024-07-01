import { useQuery } from "@tanstack/react-query";
import { getTaskList } from "../gen/rpc/task/v1/task-TaskService_connectquery"

export const useTaskList = () => {
  const taskQuery = useQuery(getTaskList.useQuery());

  return { taskQuery }
}
