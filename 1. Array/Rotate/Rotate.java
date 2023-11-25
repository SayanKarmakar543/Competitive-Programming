class Rotate{
    public static void main(String[] args) {
        int arr[] = new int[]{10, 20, 30, 40, 50};
        RotateArray(arr, 2);
        for (int i : arr) {
            System.out.print(i+" ");
        }
    }
    public static void RotateArray(int[] arr, int k){
        reverseArray(arr, 0, k-1);
        reverseArray(arr, k, arr.length-1);
        reverseArray(arr, 0, arr.length-1);

    }
    public static void reverseArray(int[] arr, int low, int high){
        while(low<high){
            int temp = arr[low];
            arr[low] = arr[high];
            arr[high] = temp;
            low++;
            high--;
        }
    }
}