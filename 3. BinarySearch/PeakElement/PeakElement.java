class PeakElement{
    public static void main(String[] args) {
        int arr[] = new int[]{1, 2, 3, 4, 5, 6, 8, 5, 1};
        System.out.println(findPeakElement(arr));      
    }
    public static int findPeakElement(int[] arr){

        // TC: O(n)
        // SC: O(1)

        for(int i=0; i<arr.length;i++){
            if(
                (i==0 || arr[i]>arr[i-1])
                    &&
                (i==arr.length-1 || arr[i] > arr[i+1])
            ){
                return arr[i];
            }
        }
        return -1;
    }
}