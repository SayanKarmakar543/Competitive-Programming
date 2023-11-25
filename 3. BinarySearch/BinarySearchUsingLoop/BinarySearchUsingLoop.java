class BinarySearchUsingLoop{
    public static void main(String[] args) {
        int arr[] = new int[]{10, 20, 30, 40, 50};
        int x = 20;
        System.out.println(BinarySearch(arr, 0, arr.length-1, x));
    }
    public static int BinarySearch(int[] arr, int low, int high, int x){
        while(low<=high){
            int mid = (low+high)/2;
            if(arr[mid]==x){
                return mid;
            }
            else if(arr[mid]>x){
                high = mid-1;
            }
            else if(arr[mid]<x){
                low = mid+1;
            }
        }
        return -1;
    }
}