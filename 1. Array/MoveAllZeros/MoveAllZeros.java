class MoveAllZeros{
    public static void main(String[] args) {
        int arr[] = new int[]{10, 20, 0, 0, 0, 30, 40, 50};
        MoveAllZerosFunc(arr);
        for (int i : arr) {
            System.out.print(i+" ");
        }
    }
    public static void MoveAllZerosFunc(int[] arr){
        // Check 0 position and set value in arr[i]
        for(int i=0;i<arr.length;i++){
            if(arr[i]==0){
                // assign the index j=i then traversing the array until get the non zero value and swap two non zero and 0 values.
                for(int j=i;j<arr.length;j++){
                    if(arr[j]!=0){
                        int temp = arr[i];
                        arr[i] = arr[j];
                        arr[j] = temp;
                        break;
                    }
                }
            }
        }
    }
}