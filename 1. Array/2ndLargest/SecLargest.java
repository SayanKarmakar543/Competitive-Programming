class SecLargest{
    public static void main(String[] args) {
        int arr[] = new int[]{60, 10, 20, 30, 40, 50, 70};
        System.out.println(FindSecLargest(arr));
    }
    public static int FindSecLargest(int arr[]){
        int largest_index = FindMaximumNumber(arr);
        int secoundLargest = arr[0];
        if(largest_index==0){
            secoundLargest = arr[1];
        }
        else if(largest_index==arr.length-1){
            secoundLargest = arr[arr.length-2];
        }
        else{
            secoundLargest = arr[0];
  
        }
        
        for(int i=0;i<arr.length;i++){
            if(arr[i]>secoundLargest && arr[i]<arr[largest_index]){
                secoundLargest = arr[i];
            }
        }
        return secoundLargest;
    }
    public static int FindMaximumNumber(int arr[]){
        int largest = arr[0];
        int largest_index = 0;
        for(int i=0;i<arr.length;i++){
            if(arr[i]>largest){
                largest = arr[i];
                largest_index = i;
            }
        }
        return largest_index;
    }
}