class MooreVoting{
    public static void main(String[] args) {
        // int arr[] = new int[]{7, 7, 5, 7, 5, 1, 5, 7, 5, 7, 5, 5, 7, 7, 5, 5, 5, 5};
        // int arr[] = new int[]{7, 7, 5, 7, 5, 1, 5, 7, 5, 7, 5, 5, 7, 7, 1, 1, 1, 1};
        int arr[] = new int[]{1, 2, 1, 1, 2};
        System.out.println(findMajorityElement(arr));
    }
    public static int findMajorityElement(int[] arr){
        int count = 0;
        int element = -1;
        for(int i=0; i<arr.length; i++){
            if(count==0 || element==arr[i]){
                element = arr[i];
                count++;
            }
            else{
                count--;
            }
        }
        count = 0;
        for(int i=0;i<arr.length; i++){
            if(element==arr[i]){
                count++;
            }
        }
        if(count>arr.length/2){
            return element;
        }
        return -1;
    }
}