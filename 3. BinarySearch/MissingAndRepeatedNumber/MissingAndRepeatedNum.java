class MissingAndRepeatedNum{
    public static void main(String[] args) {
        int arr[] = new int[]{2, 3, 2, 1, 5};
        findRepeatedMissingNum(arr);
    }
    public static void findRepeatedMissingNum(int[] arr){

        // TC: O(n^2)
        // SC: O(1)
 
		int repeating = -1;
		int missing = -1;
	    for (int i = 1; i <= arr.length; i++ ) {
		    int count = 0;
		    for (int j = 0; j < arr.length; j++) {
			    if (i == arr[j]) {
				    count++;
                }
			}
            if (count == 2) {
                repeating = i;
            }
            if (count == 0) {
                missing = i;
            }
		}
        System.out.println("Missing: "+missing);
        System.out.println("Repeating: "+repeating);
        
    }
}